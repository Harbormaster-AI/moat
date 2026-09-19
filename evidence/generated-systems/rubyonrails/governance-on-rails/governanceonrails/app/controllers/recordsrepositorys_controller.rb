class RecordsRepositorysController < ApplicationController
  def index
    @recordsRepositorys = RecordsRepository.all
  end
 
  def show
    @recordsRepository = RecordsRepository.find(params[:id])
  end
 
  def new
    @recordsRepository = RecordsRepository.new
  end
 
  def edit
    @recordsRepository = RecordsRepository.find(params[:id])
  end
 
  def create
    @recordsRepository = RecordsRepository.new(recordsRepository_params)
 
    if @recordsRepository.save
      redirect_to recordsRepositorys_path
    else
      render 'new'
    end
  end
 
  def update
    @recordsRepository = RecordsRepository.find(params[:id])
 
    if @recordsRepository.update(recordsRepository_params)
      redirect_to recordsRepositorys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @recordsRepository = RecordsRepository.find(params[:id])
    @recordsRepository.destroy
    redirect_to recordsRepositorys_path
  end

 
  private
    def recordsRepository_params
      params.require(:recordsRepository).permit(:name, :location, :ownerDepartment, :RepositoryType)
    end
end