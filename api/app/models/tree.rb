class Tree < ApplicationRecord
    belongs_to :log
    has_one :user
end
