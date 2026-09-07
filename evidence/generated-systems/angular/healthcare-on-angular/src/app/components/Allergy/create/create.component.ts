import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AllergyService } from '../../../services/Allergy.service';
import { Allergy } from '../../../models/Allergy';
import { SubBaseComponent } from '../../Allergy/sub.base.component';

@Component({
    selector: 'app-create-allergy',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAllergyComponent extends SubBaseComponent implements OnInit {

    title = 'Add Allergy';

    allergyForm: FormGroup;
    allergy: Allergy;

    constructor( http: HttpClient,
        private allergyService: AllergyService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAllergy(substance, reaction, Patient, Severity, Status): void {
        this.allergyService
        .addAllergy(substance, reaction, Patient, Severity, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAllergy']);
            });
    }

    ngOnInit(): void {
    }
}