import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BusinessUnitService } from '../../../services/BusinessUnit.service';
import { SubBaseComponent } from '../../BusinessUnit/sub.base.component';


@Component({
    selector: 'app-edit-businessUnit',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBusinessUnitComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BusinessUnit';

    businessUnitForm: FormGroup;
    businessUnit: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BusinessUnitService,
        private fb: FormBuilder
) {
        super(http);
        this.businessUnitForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      leader: ['', Validators.required],
      Organization: ['', ],
      Audits: ['', ]
        });
    }

    
    updateBusinessUnit(name, leader, Organization, Audits): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBusinessUnit(name, leader, Organization, Audits, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBusinessUnit']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBusinessUnit(params['id']).subscribe(res => {
                this.businessUnit = res;
            });
        });
    }
}