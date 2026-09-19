class WalletsController < ApplicationController
  def index
    @wallets = Wallet.all
  end
 
  def show
    @wallet = Wallet.find(params[:id])
  end
 
  def new
    @wallet = Wallet.new
  end
 
  def edit
    @wallet = Wallet.find(params[:id])
  end
 
  def create
    @wallet = Wallet.new(wallet_params)
 
    if @wallet.save
      redirect_to wallets_path
    else
      render 'new'
    end
  end
 
  def update
    @wallet = Wallet.find(params[:id])
 
    if @wallet.update(wallet_params)
      redirect_to wallets_path
    else
      render 'edit'
    end
  end
 
  def destroy
    @wallet = Wallet.find(params[:id])
    @wallet.destroy
    redirect_to wallets_path
  end

 
  private
    def wallet_params
      params.require(:wallet).permit(:currency, :balance, :Status)
    end
end