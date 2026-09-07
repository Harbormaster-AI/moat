import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PlantService } from '../../../services/Plant.service';
import { Plant } from '../../../models/Plant';
import { SubBaseComponent } from '../../Plant/sub.base.component';

@Component({
    selector: 'app-create-plant',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePlantComponent extends SubBaseComponent implements OnInit {

    title = 'Add Plant';

    plantForm: FormGroup;
    plant: Plant;

    constructor( http: HttpClient,
        private plantService: PlantService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.plantForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      plantCode: ['', Validators.required],
      address: ['', Validators.required],
      timeZone: ['', Validators.required],
      Enterprise: ['', ],
      ProductionLines: ['', ],
      WorkCenters: ['', ],
      Warehouses: ['', ],
      Assets: ['', ],
      ProductionSchedules: ['', ]
        });
    }

    
    addPlant(name, plantCode, address, timeZone, Enterprise, ProductionLines, WorkCenters, Warehouses, Assets, ProductionSchedules): void {
        this.plantService
        .addPlant(name, plantCode, address, timeZone, Enterprise, ProductionLines, WorkCenters, Warehouses, Assets, ProductionSchedules)
            .subscribe(() => {
                this.router.navigate(['/indexPlant']);
            });
    }

    ngOnInit(): void {
    }
}