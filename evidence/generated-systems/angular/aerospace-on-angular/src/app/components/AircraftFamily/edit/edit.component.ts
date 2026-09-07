import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AircraftFamilyService } from '../../../services/AircraftFamily.service';
import { SubBaseComponent } from '../../AircraftFamily/sub.base.component';


@Component({
    selector: 'app-edit-aircraftFamily',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAircraftFamilyComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AircraftFamily';

    aircraftFamilyForm: FormGroup;
    aircraftFamily: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AircraftFamilyService,
        private fb: FormBuilder
) {
        super(http);
        this.aircraftFamilyForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      familyCode: ['', Validators.required],
      Program: ['', ],
      AircraftModels: ['', ]
        });
    }

    
    updateAircraftFamily(name, familyCode, Program, AircraftModels): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAircraftFamily(name, familyCode, Program, AircraftModels, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAircraftFamily']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAircraftFamily(params['id']).subscribe(res => {
                this.aircraftFamily = res;
            });
        });
    }
}