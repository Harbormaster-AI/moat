import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { OfferService } from '../../../services/Offer.service';
import { Offer } from '../../../models/Offer';
import { SubBaseComponent } from '../../Offer/sub.base.component';

@Component({
    selector: 'app-create-offer',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateOfferComponent extends SubBaseComponent implements OnInit {

    title = 'Add Offer';

    offerForm: FormGroup;
    offer: Offer;

    constructor( http: HttpClient,
        private offerService: OfferService,
        private fb: FormBuilder,
        private router: Router
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

    
    addOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status): void {
        this.offerService
        .addOffer(offerNumber, proposedStartDate, baseSalary, signOnBonus, Requisition, Candidate, ApprovedBy, Contract, Status)
            .subscribe(() => {
                this.router.navigate(['/indexOffer']);
            });
    }

    ngOnInit(): void {
    }
}