
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAudienceSegmentComponent } from './index.component';
import { AudienceSegmentService } from '../../../services/AudienceSegment.service';

describe('IndexAudienceSegmentComponent', () => {
  let component: IndexAudienceSegmentComponent;
  let fixture: ComponentFixture<IndexAudienceSegmentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAudienceSegmentComponent
      ],
      providers: [
        AudienceSegmentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAudienceSegmentComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});