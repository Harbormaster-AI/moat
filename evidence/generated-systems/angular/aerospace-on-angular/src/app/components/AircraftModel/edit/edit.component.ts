import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftModelService } from '../../../services/AircraftModel.service';
import { SubBaseComponent } from '../../AircraftModel/sub.base.component';


@Component({
    selector: 'app-edit-aircraftModel',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftModelComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftModel';

    aircraftModelForm: FormGroup;
    aircraftModel: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftModelService,
        private fb: FormBuilder
) {
        super(http);
        this.aircraftModelForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      modelDesignation: ['', Validators.required],
      Family: ['', ],
      Variants: ['', ],
      EngineTypes: ['', ],
      AircraftType: ['', ]
        });
    }

    
    updateAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftModel(name, modelDesignation, Family, Variants, EngineTypes, AircraftType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftModel']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftModel(params['id']).subscribe(res => {
                this.aircraftModel = res;
            });
        });
    }
}