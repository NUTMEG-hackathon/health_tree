class CreateLogs < ActiveRecord::Migration[6.1]
  def change
    create_table :logs do |t|
      t.boolean :win

      t.timestamps
    end
  end
end
