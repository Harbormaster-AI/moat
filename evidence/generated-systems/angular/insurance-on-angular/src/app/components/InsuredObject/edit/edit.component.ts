import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { InsuredObjectService } from '../../../services/InsuredObject.service';
import { SubBaseComponent } from '../../InsuredObject/sub.base.component';


@Component({
    selector: 'app-edit-insuredObject',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditInsuredObjectComponent extends SubBaseComponent implements OnInit {

    title = 'Edit InsuredObject';

    insuredObjectForm: FormGroup;
    insuredObject: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: InsuredObjectService,
        private fb: FormBuilder
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

    
    updateInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateInsuredObject(description, serialOrId, primaryAddress, Policy, Coverages, ObjectType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexInsuredObject']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getInsuredObject(params['id']).subscribe(res => {
                this.insuredObject = res;
            });
        });
    }
}