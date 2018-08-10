require './lib/card.rb'
require './lib/deck.rb'
require './lib/player.rb'

def start
  # Open a fresh pack and shuffle it
  @d = Deck.new
  @d.deck.shuffle!

  # Get someone to play
  p1 = Player.new
  p2 = Player.new
  
  # Deal hands out
  deal_hand(p1, 7)
  puts "Player 1 hand:"
  p1.hand.each do |card|
    puts card.friendly_name
  end

  deal_hand(p2, 7)
  puts "Player 2 hand:"
  p2.hand.each do |card|
    puts card.friendly_name
  end
end

def deal_hand(player, number_of_cards)
  zero_based_number_of_cards = number_of_cards - 1
  hand_dealt = @d.deck.take(number_of_cards)
  @d.deck.slice!(0..zero_based_number_of_cards)
  player.set_hand(hand_dealt)
end

def play_trick(player, index_of_card_in_hand)
  # Take a card from the player's hand and play a trick
end

start