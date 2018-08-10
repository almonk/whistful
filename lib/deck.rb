class Deck
  def initialize
    @deck = []
    suits = ["Hearts", "Spades", "Clubs", "Diamonds"]
    currentSuit = 0
    
    # Initialise a full deck
    while currentSuit < 4
      i = 1
      while i < 14
        c = Card.new(i, suits[currentSuit])
        @deck << c
        i += 1
      end
      currentSuit += 1
    end
  end

  def deck
    @deck
  end

  def shuffle
    @deck.shuffle!
  end
end