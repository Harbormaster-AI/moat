
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProductionScheduleComponent } from './create.component';
import { ProductionScheduleService } from '../../../services/ProductionSchedule.service';
import { Router } from '@angular/router';

describe('CreateProductionScheduleComponent', () => {
  let component: CreateProductionScheduleComponent;
  let fixture: ComponentFixture<CreateProductionScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProductionScheduleComponent
      ],
      providers: [
        ProductionScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProductionScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});