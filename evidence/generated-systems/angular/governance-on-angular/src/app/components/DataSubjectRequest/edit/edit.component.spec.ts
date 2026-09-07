
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditDataSubjectRequestComponent } from './edit.component';
import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';

describe('EditDataSubjectRequestComponent', () => {
  let component: EditDataSubjectRequestComponent;
  let fixture: ComponentFixture<EditDataSubjectRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditDataSubjectRequestComponent
      ],
      providers: [
        DataSubjectRequestService,
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

    fixture = TestBed.createComponent(EditDataSubjectRequestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});