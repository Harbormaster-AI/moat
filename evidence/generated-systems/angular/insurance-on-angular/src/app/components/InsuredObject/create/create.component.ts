import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InsuredObjectService } from '../../../services/InsuredObject.service';
import { InsuredObject } from '../../../models/InsuredObject';
import { SubBaseComponent } from '../../InsuredObject/sub.base.component';

@Component({
    selector: 'app-create-insuredObject',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInsuredObjectComponent extends SubBaseComponent implements OnInit {

    title = 'Add InsuredObject';

    insuredObjectForm: FormGroup;
    insuredObject: InsuredObject;

    constructor( http: HttpClient,
        private insuredObjectService: InsuredObjectService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.insuredObjectForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  description: ['', Validators.required],
      serialOrId: ['', Validators.required],
      primaryAddress: ['', Validators.required],
      Policy: ['', ],
      Coverages: ['', ],
      ObjectType: ['', ]
        });
    }

    
    addInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType): void {
        this.insuredObjectService
        .addInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType)
            .subscribe(() => {
                this.router.navigate(['/indexInsuredObject']);
            });
    }

    ngOnInit(): void {
    }
}