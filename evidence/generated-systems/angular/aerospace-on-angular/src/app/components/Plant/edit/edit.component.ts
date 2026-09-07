import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PlantService } from '../../../services/Plant.service';
import { SubBaseComponent } from '../../Plant/sub.base.component';


@Component({
    selector: 'app-edit-plant',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPlantComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Plant';

    plantForm: FormGroup;
    plant: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PlantService,
        private fb: FormBuilder
) {
        super(http);
        this.plantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      plantCode: ['', Validators.required],
      address: ['', Validators.required],
      Manufacturer: ['', ],
      ProductionLines: ['', ],
      Warehouses: ['', ]
        });
    }

    
    updatePlant(name, plantCode, address, Manufacturer, ProductionLines, Warehouses): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePlant(name, plantCode, address, Manufacturer, ProductionLines, Warehouses, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPlant']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPlant(params['id']).subscribe(res => {
                this.plant = res;
            });
        });
    }
}