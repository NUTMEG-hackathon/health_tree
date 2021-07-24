class CreateAnswers < ActiveRecord::Migration[6.1]
  def change
    create_table :answers do |t|
      t.string :title
      t.integer :point
      t.integer :question_id

      t.timestamps
    end
  end
end
