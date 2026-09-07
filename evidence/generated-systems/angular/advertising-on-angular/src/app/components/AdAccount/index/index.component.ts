
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AdAccountService } from '../../../services/AdAccount.service';
import { AdAccount } from '../../../models/AdAccount';

@Component({
    selector: 'app-index-adAccount',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAdAccountComponent implements OnInit {

    adAccounts: AdAccount[] = [];

    constructor(
        private router: Router,
        private service: AdAccountService
) {}

    ngOnInit(): void {
        this.getAdAccounts();
}

    getAdAccounts(): void {
        this.service.getAdAccounts().subscribe((res) => {
        this.adAccounts = res;
    });
}

    deleteAdAccount(id: any): void {
        this.service.deleteAdAccount(id)
            .subscribe(() => {
                this.getAdAccounts();
            });
    }
}