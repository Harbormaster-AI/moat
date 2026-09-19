class FraudScenariosController < ApplicationController
  def index
    @fraudScenarios = FraudScenario.all
  end
 
  def show
    @fraudScenario = FraudScenario.find(params[:id])
  end
 
  def new
    @fraudScenario = FraudScenario.new
  end
 
  def edit
    @fraudScenario = FraudScenario.find(params[:id])
  end
 
  def create
    @fraudScenario = FraudScenario.new(fraudScenario_params)
 
    if @fraudScenario.save
      redirect_to fraudScenarios_path
    else
      render 'new'
    end
  end
 
  def update
    @fraudScenario = FraudScenario.find(params[:id])
 
    if @fraudScenario.update(fraudScenario_params)
      redirect_to fraudScenarios_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @fraudScenario = FraudScenario.find(params[:id])
    @fraudScenario.destroy
    redirect_to fraudScenarios_path
  end

 
  private
    def fraudScenario_params
      params.require(:fraudScenario).permit(:name, :riskAppetite, :DetectionType)
    end
end