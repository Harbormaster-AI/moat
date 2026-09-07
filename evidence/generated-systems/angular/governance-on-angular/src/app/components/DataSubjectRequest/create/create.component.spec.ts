
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataSubjectRequestComponent } from './create.component';
import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';
import { Router } from '@angular/router';

describe('CreateDataSubjectRequestComponent', () => {
  let component: CreateDataSubjectRequestComponent;
  let fixture: ComponentFixture<CreateDataSubjectRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataSubjectRequestComponent
      ],
      providers: [
        DataSubjectRequestService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataSubjectRequestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});