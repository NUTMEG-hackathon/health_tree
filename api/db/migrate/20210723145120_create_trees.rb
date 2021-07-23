class CreateTrees < ActiveRecord::Migration[6.1]
  def change
    create_table :trees do |t|
      t.string :name
      t.float :point
      t.integer :log_id

      t.timestamps
    end
  end
end
