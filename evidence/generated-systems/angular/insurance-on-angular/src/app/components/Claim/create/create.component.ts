import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ClaimService } from '../../../services/Claim.service';
import { Claim } from '../../../models/Claim';
import { SubBaseComponent } from '../../Claim/sub.base.component';

@Component({
    selector: 'app-create-claim',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateClaimComponent extends SubBaseComponent implements OnInit {

    title = 'Add Claim';

    claimForm: FormGroup;
    claim: Claim;

    constructor( http: HttpClient,
        private claimService: ClaimService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.claimForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  claimNumber: ['', Validators.required],
      noticeDate: ['', Validators.required],
      lossDate: ['', Validators.required],
      reportedBy: ['', Validators.required],
      Policy: ['', ],
      Customer: ['', ],
      Adjuster: ['', ],
      Incident: ['', ],
      Exposures: ['', ],
      Reserves: ['', ],
      ClaimPayments: ['', ],
      ServiceProviders: ['', ],
      Subrogations: ['', ],
      Status: ['', ],
      LossCause: ['', ]
        });
    }

    
    addClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause): void {
        this.claimService
        .addClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause)
            .subscribe(() => {
                this.router.navigate(['/indexClaim']);
            });
    }

    ngOnInit(): void {
    }
}