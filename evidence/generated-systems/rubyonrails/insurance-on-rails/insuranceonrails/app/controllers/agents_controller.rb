class AgentsController < ApplicationController
  def index
    @agents = Agent.all
  end
 
  def show
    @agent = Agent.find(params[:id])
  end
 
  def new
    @agent = Agent.new
  end
 
  def edit
    @agent = Agent.find(params[:id])
  end
 
  def create
    @agent = Agent.new(agent_params)
 
    if @agent.save
      redirect_to agents_path
    else
      render 'new'
    end
  end
 
  def update
    @agent = Agent.find(params[:id])
 
    if @agent.update(agent_params)
      redirect_to agents_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @agent = Agent.find(params[:id])
    @agent.destroy
    redirect_to agents_path
  end

 
  private
    def agent_params
      params.require(:agent).permit(:firstName, :lastName, :licenseId, :Status)
    end
end