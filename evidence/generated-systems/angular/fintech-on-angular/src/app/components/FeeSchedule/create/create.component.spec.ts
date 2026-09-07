
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateFeeScheduleComponent } from './create.component';
import { FeeScheduleService } from '../../../services/FeeSchedule.service';
import { Router } from '@angular/router';

describe('CreateFeeScheduleComponent', () => {
  let component: CreateFeeScheduleComponent;
  let fixture: ComponentFixture<CreateFeeScheduleComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateFeeScheduleComponent
      ],
      providers: [
        FeeScheduleService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateFeeScheduleComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});