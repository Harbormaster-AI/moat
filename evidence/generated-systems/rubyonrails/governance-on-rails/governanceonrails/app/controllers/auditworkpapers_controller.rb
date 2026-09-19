class AuditWorkpapersController < ApplicationController
  def index
    @auditWorkpapers = AuditWorkpaper.all
  end
 
  def show
    @auditWorkpaper = AuditWorkpaper.find(params[:id])
  end
 
  def new
    @auditWorkpaper = AuditWorkpaper.new
  end
 
  def edit
    @auditWorkpaper = AuditWorkpaper.find(params[:id])
  end
 
  def create
    @auditWorkpaper = AuditWorkpaper.new(auditWorkpaper_params)
 
    if @auditWorkpaper.save
      redirect_to auditWorkpapers_path
    else
      render 'new'
    end
  end
 
  def update
    @auditWorkpaper = AuditWorkpaper.find(params[:id])
 
    if @auditWorkpaper.update(auditWorkpaper_params)
      redirect_to auditWorkpapers_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @auditWorkpaper = AuditWorkpaper.find(params[:id])
    @auditWorkpaper.destroy
    redirect_to auditWorkpapers_path
  end

 
  private
    def auditWorkpaper_params
      params.require(:auditWorkpaper).permit(:workpaperRef, :subject, :workpaperUrl)
    end
end