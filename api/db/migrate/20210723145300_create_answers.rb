class CreateAnswers < ActiveRecord::Migration[6.1]
  def change
    create_table :answers do |t|
      t.integer :title
      t.float :point

      t.timestamps
    end
  end
end
