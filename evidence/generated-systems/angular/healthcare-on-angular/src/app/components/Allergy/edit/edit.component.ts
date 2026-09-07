import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AllergyService } from '../../../services/Allergy.service';
import { SubBaseComponent } from '../../Allergy/sub.base.component';


@Component({
    selector: 'app-edit-allergy',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAllergyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Allergy';

    allergyForm: FormGroup;
    allergy: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AllergyService,
        private fb: FormBuilder
) {
        super(http);
        this.allergyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  substance: ['', Validators.required],
      reaction: ['', Validators.required],
      Patient: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    updateAllergy(substance, reaction, Patient, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAllergy(substance, reaction, Patient, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAllergy']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAllergy(params['id']).subscribe(res => {
                this.allergy = res;
            });
        });
    }
}