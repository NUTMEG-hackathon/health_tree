class CreateQuestions < ActiveRecord::Migration[6.1]
  def change
    create_table :questions do |t|
      t.integer :title
      t.integer :answer_id

      t.timestamps
    end
  end
end
