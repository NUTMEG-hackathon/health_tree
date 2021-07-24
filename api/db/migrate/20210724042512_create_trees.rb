class CreateTrees < ActiveRecord::Migration[6.1]
  def change
    create_table :trees do |t|
      t.string :name
      t.float :point
      t.integer :user_id
      t.integer :type

      t.timestamps
    end
  end
end
