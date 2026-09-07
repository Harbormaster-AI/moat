
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateWorkShiftComponent } from './create.component';
import { WorkShiftService } from '../../../services/WorkShift.service';
import { Router } from '@angular/router';

describe('CreateWorkShiftComponent', () => {
  let component: CreateWorkShiftComponent;
  let fixture: ComponentFixture<CreateWorkShiftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateWorkShiftComponent
      ],
      providers: [
        WorkShiftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateWorkShiftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});