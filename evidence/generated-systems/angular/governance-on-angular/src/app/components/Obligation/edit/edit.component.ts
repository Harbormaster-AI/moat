import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ObligationService } from '../../../services/Obligation.service';
import { SubBaseComponent } from '../../Obligation/sub.base.component';


@Component({
    selector: 'app-edit-obligation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditObligationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Obligation';

    obligationForm: FormGroup;
    obligation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ObligationService,
        private fb: FormBuilder
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

    
    updateObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency): void {
        this.route.params.subscribe((params) => {

                        this.service.updateObligation(referenceNumber, descriptionText, Regulation, Controls, Policies, Contracts, ObligationType, ReviewFrequency, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexObligation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getObligation(params['id']).subscribe(res => {
                this.obligation = res;
            });
        });
    }
}