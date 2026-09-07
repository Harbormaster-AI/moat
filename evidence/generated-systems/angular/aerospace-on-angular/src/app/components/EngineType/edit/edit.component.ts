import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EngineTypeService } from '../../../services/EngineType.service';
import { SubBaseComponent } from '../../EngineType/sub.base.component';


@Component({
    selector: 'app-edit-engineType',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEngineTypeComponent extends SubBaseComponent implements OnInit {

    title = 'Edit EngineType';

    engineTypeForm: FormGroup;
    engineType: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EngineTypeService,
        private fb: FormBuilder
) {
        super(http);
        this.engineTypeForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  engineModelCode: ['', Validators.required],
      maxThrustKn: ['', Validators.required],
      Supplier: ['', ],
      CompatibleModels: ['', ],
      Category: ['', ]
        });
    }

    
    updateEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEngineType(engineModelCode, maxThrustKn, Supplier, CompatibleModels, Category, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEngineType']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEngineType(params['id']).subscribe(res => {
                this.engineType = res;
            });
        });
    }
}