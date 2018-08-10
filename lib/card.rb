class Card
  def initialize(rank, suit)
    @rank = rank
    @suit = suit
  end

  def rank
    @rank
  end

  def suit
    @suit
  end

  def friendly_name
    friendly_name = ""
    case @rank
    when 11
      friendly_name = "Jack"
    when 12
      friendly_name = "Queen"
    when 13
      friendly_name = "King"
    when 14
      friendly_name = "ace"
    else
      friendly_name = @rank.to_s
    end

    "The #{friendly_name} of #{self.suit}"
  end
end