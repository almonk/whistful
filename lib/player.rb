class Player
  def initialize(number)
    @hand = []
    @number = number
  end

  def set_hand(array)
    @hand = array
  end

  def hand
    @hand
  end

  def number
    @number
  end
end