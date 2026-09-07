
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditMaintenancePlanComponent } from './edit.component';
import { MaintenancePlanService } from '../../../services/MaintenancePlan.service';

describe('EditMaintenancePlanComponent', () => {
  let component: EditMaintenancePlanComponent;
  let fixture: ComponentFixture<EditMaintenancePlanComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditMaintenancePlanComponent
      ],
      providers: [
        MaintenancePlanService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditMaintenancePlanComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});