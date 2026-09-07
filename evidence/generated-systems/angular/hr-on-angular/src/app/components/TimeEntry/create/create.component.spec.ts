
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTimeEntryComponent } from './create.component';
import { TimeEntryService } from '../../../services/TimeEntry.service';
import { Router } from '@angular/router';

describe('CreateTimeEntryComponent', () => {
  let component: CreateTimeEntryComponent;
  let fixture: ComponentFixture<CreateTimeEntryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTimeEntryComponent
      ],
      providers: [
        TimeEntryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTimeEntryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});