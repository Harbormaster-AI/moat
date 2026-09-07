import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ClaimReserveService } from '../../../services/ClaimReserve.service';
import { ClaimReserve } from '../../../models/ClaimReserve';
import { SubBaseComponent } from '../../ClaimReserve/sub.base.component';

@Component({
    selector: 'app-create-claimReserve',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateClaimReserveComponent extends SubBaseComponent implements OnInit {

    title = 'Add ClaimReserve';

    claimReserveForm: FormGroup;
    claimReserve: ClaimReserve;

    constructor( http: HttpClient,
        private claimReserveService: ClaimReserveService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.claimReserveForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  amount: ['', Validators.required],
      setDate: ['', Validators.required],
      Claim: ['', ],
      Exposure: ['', ],
      ReserveType: ['', ],
      Status: ['', ]
        });
    }

    
    addClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status): void {
        this.claimReserveService
        .addClaimReserve(amount, setDate, Claim, Exposure, ReserveType, Status)
            .subscribe(() => {
                this.router.navigate(['/indexClaimReserve']);
            });
    }

    ngOnInit(): void {
    }
}