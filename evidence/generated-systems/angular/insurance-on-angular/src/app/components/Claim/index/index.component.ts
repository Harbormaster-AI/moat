
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ClaimService } from '../../../services/Claim.service';
import { Claim } from '../../../models/Claim';

@Component({
    selector: 'app-index-claim',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexClaimComponent implements OnInit {

    claims: Claim[] = [];

    constructor(
        private router: Router,
        private service: ClaimService
) {}

    ngOnInit(): void {
        this.getClaims();
}

    getClaims(): void {
        this.service.getClaims().subscribe((res) => {
        this.claims = res;
    });
}

    deleteClaim(id: any): void {
        this.service.deleteClaim(id)
            .subscribe(() => {
                this.getClaims();
            });
    }
}