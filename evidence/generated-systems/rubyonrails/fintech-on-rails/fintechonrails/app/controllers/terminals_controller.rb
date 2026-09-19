class TerminalsController < ApplicationController
  def index
    @terminals = Terminal.all
  end
 
  def show
    @terminal = Terminal.find(params[:id])
  end
 
  def new
    @terminal = Terminal.new
  end
 
  def edit
    @terminal = Terminal.find(params[:id])
  end
 
  def create
    @terminal = Terminal.new(terminal_params)
 
    if @terminal.save
      redirect_to terminals_path
    else
      render 'new'
    end
  end
 
  def update
    @terminal = Terminal.find(params[:id])
 
    if @terminal.update(terminal_params)
      redirect_to terminals_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @terminal = Terminal.find(params[:id])
    @terminal.destroy
    redirect_to terminals_path
  end

 
  private
    def terminal_params
      params.require(:terminal).permit(:location, :Type, :Status)
    end
end