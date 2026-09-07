
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAudienceSegmentComponent } from './create.component';
import { AudienceSegmentService } from '../../../services/AudienceSegment.service';
import { Router } from '@angular/router';

describe('CreateAudienceSegmentComponent', () => {
  let component: CreateAudienceSegmentComponent;
  let fixture: ComponentFixture<CreateAudienceSegmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAudienceSegmentComponent
      ],
      providers: [
        AudienceSegmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAudienceSegmentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});