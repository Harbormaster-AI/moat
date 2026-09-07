
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditScheduleExceptionComponent } from './edit.component';
import { ScheduleExceptionService } from '../../../services/ScheduleException.service';

describe('EditScheduleExceptionComponent', () => {
  let component: EditScheduleExceptionComponent;
  let fixture: ComponentFixture<EditScheduleExceptionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditScheduleExceptionComponent
      ],
      providers: [
        ScheduleExceptionService,
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

    fixture = TestBed.createComponent(EditScheduleExceptionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});