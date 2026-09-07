import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ClaimService } from '../../../services/Claim.service';
import { SubBaseComponent } from '../../Claim/sub.base.component';


@Component({
    selector: 'app-edit-claim',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditClaimComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Claim';

    claimForm: FormGroup;
    claim: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ClaimService,
        private fb: FormBuilder
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

    
    updateClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause): void {
        this.route.params.subscribe((params) => {

                        this.service.updateClaim(claimNumber, noticeDate, lossDate, reportedBy, Policy, Customer, Adjuster, Incident, Exposures, Reserves, ClaimPayments, ServiceProviders, Subrogations, Status, LossCause, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexClaim']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getClaim(params['id']).subscribe(res => {
                this.claim = res;
            });
        });
    }
}