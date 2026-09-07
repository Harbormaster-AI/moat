
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDeviceCriterionComponent } from './create.component';
import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';
import { Router } from '@angular/router';

describe('CreateDeviceCriterionComponent', () => {
  let component: CreateDeviceCriterionComponent;
  let fixture: ComponentFixture<CreateDeviceCriterionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDeviceCriterionComponent
      ],
      providers: [
        DeviceCriterionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDeviceCriterionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});