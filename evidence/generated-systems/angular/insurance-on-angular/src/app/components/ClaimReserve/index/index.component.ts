
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ClaimReserveService } from '../../../services/ClaimReserve.service';
import { ClaimReserve } from '../../../models/ClaimReserve';

@Component({
    selector: 'app-index-claimReserve',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexClaimReserveComponent implements OnInit {

    claimReserves: ClaimReserve[] = [];

    constructor(
        private router: Router,
        private service: ClaimReserveService
) {}

    ngOnInit(): void {
        this.getClaimReserves();
}

    getClaimReserves(): void {
        this.service.getClaimReserves().subscribe((res) => {
        this.claimReserves = res;
    });
}

    deleteClaimReserve(id: any): void {
        this.service.deleteClaimReserve(id)
            .subscribe(() => {
                this.getClaimReserves();
            });
    }
}