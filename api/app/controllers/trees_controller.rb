class TreesController < ApplicationController
  before_action :set_tree, only: %i[ show update destroy ]

  # GET /trees
  # GET /trees.json
  def index
    @trees = Tree.all
  end

  # GET /trees/1
  # GET /trees/1.json
  def show
  end

  # POST /trees
  # POST /trees.json
  def create
    @tree = Tree.new(tree_params)

    if @tree.save
      render :show, status: :created, location: @tree
    else
      render json: @tree.errors, status: :unprocessable_entity
    end
  end

  # PATCH/PUT /trees/1
  # PATCH/PUT /trees/1.json
  def update
    if @tree.update(tree_params)
      render :show, status: :ok, location: @tree
    else
      render json: @tree.errors, status: :unprocessable_entity
    end
  end

  # DELETE /trees/1
  # DELETE /trees/1.json
  def destroy
    @tree.destroy
  end

  private
    # Use callbacks to share common setup or constraints between actions.
    def set_tree
      @tree = Tree.find(params[:id])
    end

    # Only allow a list of trusted parameters through.
    def tree_params
      params.require(:tree).permit(:name, :point, :user_id, :type)
    end
end
