require './lib/card.rb'
require './lib/deck.rb'
require './lib/player.rb'

def start
  # Open a fresh pack and shuffle it
  @d = Deck.new
  @d.deck.shuffle!
  @current_trick = [] # The current leading trick of the round e.g. 10 of Hearts
  @current_trump_suit = pick_trumps

  # Get someone to play
  @p1 = Player.new(1)
  @p2 = Player.new(2)
  
  # Deal hands out
  deal_hand(@p1, 7)
  deal_hand(@p2, 7)
  
  ask_player_for_trick(@p1)
  ask_cpu_player_to_respond_to_trick(@p2)
end

def deal_hand(player, number_of_cards)
  zero_based_number_of_cards = number_of_cards - 1
  hand_dealt = @d.deck.take(number_of_cards)
  @d.deck.slice!(0..zero_based_number_of_cards)
  player.set_hand(hand_dealt)
end

def list_hand(player)
  puts "==========================="
  puts "Player #{player.number}'s hand"
  i = 0
  player.hand.each do |card|
    puts "[#{i}] #{card.friendly_name}"
    i+=1
  end
  puts "==========================="
end

def pick_trumps
  @d.deck.first.suit
  puts "Trumps is #{@d.deck.first.suit}"
end

def ask_player_for_trick(player)
  list_hand(@p1)

  puts "Which trick would you like to play?"
  index_of_card_in_hand = gets.to_i

  unless index_of_card_in_hand > 6
    play_trick(player, index_of_card_in_hand)
  else
    # Invalid out of bounds
    ask_player_for_trick(player)
  end
end

def play_trick(player, index_of_card_in_hand)
  # Take a card from the player's hand and play a trick
  played_card = player.hand[index_of_card_in_hand]
  player.hand.slice!(index_of_card_in_hand)
  @current_trick << played_card
  puts "Played #{played_card.friendly_name}"
end

def can_pick_up_from_deck?
  # There are cards left in the deck that can be picked up
  if @d.deck.count > 2
    true
  else
    false
  end
end

def ask_cpu_player_to_play_trick(cpu_player)
  # CPU player won the last round and goes first...
end

def ask_cpu_player_to_respond_to_trick(cpu_player)
  puts "CPU is thinking..."
  sleep(1)
  matching_suit = cpu_player.hand.select { |card| card.suit == @current_trick.first.suit }
  
  if matching_suit.count > 0
    # The CPU player has cards that match the played trick
    # and so must play them...
    # Now to find out if the CPU can beat the current trick...
    # Lets sort the cards low -> high
    sorted_matching_suit = matching_suit.sort_by { |k| k.rank }
    can_outrank = false
    index_of_first_outranking_card = 0
    i = 0
    sorted_matching_suit.each do |card|
      if card.rank > @current_trick.first.rank
        can_outrank = true
        index_of_first_outranking_card = i
      else
        can_outrank = false
      end
      i += 1
    end

    if can_outrank
      # Play the outranking card
      play_trick(cpu_player, cpu_player.hand.index(sorted_matching_suit[index_of_first_outranking_card]))
      round_is_won_by(cpu_player)
    else
      play_trick(cpu_player, cpu_player.hand.index(sorted_matching_suit.first))
      round_is_won_by(@p1)
    end

  else
    # The CPU must attempt to trump or discard...
    puts "I have no #{@current_trick.first.suit}"
    @trump_cards = cpu_player.hand.select { |card| card.suit == @current_trump_suit }
    if @trump_cards.count > 0
      # The CPU player has a trump card !
      puts "CPU can trump..."
    else
      puts "CPU must discard..."
      # Discard lowest ranking card, preferably not a trump
    end
  end
end

def round_is_won_by(player)
  if player.number == 1
    puts "You win the round!"
    ask_player_for_trick(@p1)
    ask_cpu_player_to_respond_to_trick(@p2)
  elsif player.number == 2
    puts "CPU won the round."
  end
end

start