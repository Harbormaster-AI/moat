import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ObligationService } from '../../../services/Obligation.service';
import { Obligation } from '../../../models/Obligation';
import { SubBaseComponent } from '../../Obligation/sub.base.component';

@Component({
    selector: 'app-create-obligation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateObligationComponent extends SubBaseComponent implements OnInit {

    title = 'Add Obligation';

    obligationForm: FormGroup;
    obligation: Obligation;

    constructor( http: HttpClient,
        private obligationService: ObligationService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.obligationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  referenceNumber: ['', Validators.required],
      descriptionText: ['', Validators.required],
      Regulation: ['', ],
      Controls: ['', ],
      Policies: ['', ],
      Contracts: ['', ],
      ObligationType: ['', ],
      ReviewFrequency: ['', ]
        });
    }

    
    addObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency): void {
        this.obligationService
        .addObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency)
            .subscribe(() => {
                this.router.navigate(['/indexObligation']);
            });
    }

    ngOnInit(): void {
    }
}