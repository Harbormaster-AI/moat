import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EndorsementService } from '../../../services/Endorsement.service';
import { SubBaseComponent } from '../../Endorsement/sub.base.component';


@Component({
    selector: 'app-edit-endorsement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEndorsementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Endorsement';

    endorsementForm: FormGroup;
    endorsement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EndorsementService,
        private fb: FormBuilder
) {
        super(http);
        this.endorsementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  endorsementNumber: ['', Validators.required],
      effectiveDate: ['', Validators.required],
      description: ['', Validators.required],
      Policy: ['', ]
        });
    }

    
    updateEndorsement(endorsementNumber, effectiveDate, description, Policy): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEndorsement(endorsementNumber, effectiveDate, description, Policy, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEndorsement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEndorsement(params['id']).subscribe(res => {
                this.endorsement = res;
            });
        });
    }
}