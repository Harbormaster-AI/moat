class DataSourcesController < ApplicationController
  def index
    @dataSources = DataSource.all
  end
 
  def show
    @dataSource = DataSource.find(params[:id])
  end
 
  def new
    @dataSource = DataSource.new
  end
 
  def edit
    @dataSource = DataSource.find(params[:id])
  end
 
  def create
    @dataSource = DataSource.new(dataSource_params)
 
    if @dataSource.save
      redirect_to dataSources_path
    else
      render 'new'
    end
  end
 
  def update
    @dataSource = DataSource.find(params[:id])
 
    if @dataSource.update(dataSource_params)
      redirect_to dataSources_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @dataSource = DataSource.find(params[:id])
    @dataSource.destroy
    redirect_to dataSources_path
  end

 
  private
    def dataSource_params
      params.require(:dataSource).permit(:name, :connection, :Streaming, :SourceType, :Format)
    end
end