class LineageNodesController < ApplicationController
  def index
    @lineageNodes = LineageNode.all
  end
 
  def show
    @lineageNode = LineageNode.find(params[:id])
  end
 
  def new
    @lineageNode = LineageNode.new
  end
 
  def edit
    @lineageNode = LineageNode.find(params[:id])
  end
 
  def create
    @lineageNode = LineageNode.new(lineageNode_params)
 
    if @lineageNode.save
      redirect_to lineageNodes_path
    else
      render 'new'
    end
  end
 
  def update
    @lineageNode = LineageNode.find(params[:id])
 
    if @lineageNode.update(lineageNode_params)
      redirect_to lineageNodes_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @lineageNode = LineageNode.find(params[:id])
    @lineageNode.destroy
    redirect_to lineageNodes_path
  end

 
  private
    def lineageNode_params
      params.require(:lineageNode).permit(:name, :qualifiedName, :NodeType)
    end
end