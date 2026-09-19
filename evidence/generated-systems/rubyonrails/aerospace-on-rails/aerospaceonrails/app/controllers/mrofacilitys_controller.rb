class MROFacilitysController < ApplicationController
  def index
    @mROFacilitys = MROFacility.all
  end
 
  def show
    @mROFacility = MROFacility.find(params[:id])
  end
 
  def new
    @mROFacility = MROFacility.new
  end
 
  def edit
    @mROFacility = MROFacility.find(params[:id])
  end
 
  def create
    @mROFacility = MROFacility.new(mROFacility_params)
 
    if @mROFacility.save
      redirect_to mROFacilitys_path
    else
      render 'new'
    end
  end
 
  def update
    @mROFacility = MROFacility.find(params[:id])
 
    if @mROFacility.update(mROFacility_params)
      redirect_to mROFacilitys_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @mROFacility = MROFacility.find(params[:id])
    @mROFacility.destroy
    redirect_to mROFacilitys_path
  end

 
  private
    def mROFacility_params
      params.require(:mROFacility).permit(:name, :approvalScope, :address)
    end
end