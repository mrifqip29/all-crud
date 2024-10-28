class AddCompletedToTodo < ActiveRecord::Migration[7.2]
  def change
    add_column :todos, :completed, :boolean, default: false
  end
end
