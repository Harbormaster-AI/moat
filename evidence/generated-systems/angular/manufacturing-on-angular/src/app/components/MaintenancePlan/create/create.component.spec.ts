
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateMaintenancePlanComponent } from './create.component';
import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';
import { Router } from '@angular/router';

describe('CreateMaintenancePlanComponent', () => {
  let component: CreateMaintenancePlanComponent;
  let fixture: ComponentFixture<CreateMaintenancePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateMaintenancePlanComponent
      ],
      providers: [
        MaintenancePlanService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateMaintenancePlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});