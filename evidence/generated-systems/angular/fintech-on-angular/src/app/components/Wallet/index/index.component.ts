
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WalletService } from '../../../services/Wallet.service';
import { Wallet } from '../../../models/Wallet';

@Component({
    selector: 'app-index-wallet',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWalletComponent implements OnInit {

    wallets: Wallet[] = [];

    constructor(
        private router: Router,
        private service: WalletService
) {}

    ngOnInit(): void {
        this.getWallets();
}

    getWallets(): void {
        this.service.getWallets().subscribe((res) => {
        this.wallets = res;
    });
}

    deleteWallet(id: any): void {
        this.service.deleteWallet(id)
            .subscribe(() => {
                this.getWallets();
            });
    }
}