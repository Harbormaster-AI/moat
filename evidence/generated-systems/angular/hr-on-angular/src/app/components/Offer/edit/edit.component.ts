import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { OfferService } from '../../../services/Offer.service';
import { SubBaseComponent } from '../../Offer/sub.base.component';


@Component({
    selector: 'app-edit-offer',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditOfferComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Offer';

    offerForm: FormGroup;
    offer: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: OfferService,
        private fb: FormBuilder
) {
        super(http);
        this.offerForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  offerNumber: ['', Validators.required],
      proposedStartDate: ['', Validators.required],
      baseSalary: ['', Validators.required],
      signOnBonus: ['', Validators.required],
      Requisition: ['', ],
      Candidate: ['', ],
      ApprovedBy: ['', ],
      Contract: ['', ],
      Status: ['', ]
        });
    }

    
    updateOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexOffer']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getOffer(params['id']).subscribe(res => {
                this.offer = res;
            });
        });
    }
}