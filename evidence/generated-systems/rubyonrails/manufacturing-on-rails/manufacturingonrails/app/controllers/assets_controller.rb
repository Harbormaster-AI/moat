class AssetsController < ApplicationController
  def index
    @assets = Asset.all
  end
 
  def show
    @asset = Asset.find(params[:id])
  end
 
  def new
    @asset = Asset.new
  end
 
  def edit
    @asset = Asset.find(params[:id])
  end
 
  def create
    @asset = Asset.new(asset_params)
 
    if @asset.save
      redirect_to assets_path
    else
      render 'new'
    end
  end
 
  def update
    @asset = Asset.find(params[:id])
 
    if @asset.update(asset_params)
      redirect_to assets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @asset = Asset.find(params[:id])
    @asset.destroy
    redirect_to assets_path
  end

 
  private
    def asset_params
      params.require(:asset).permit(:assetTag, :assetName, :commissioningDate, :AssetStatus)
    end
end