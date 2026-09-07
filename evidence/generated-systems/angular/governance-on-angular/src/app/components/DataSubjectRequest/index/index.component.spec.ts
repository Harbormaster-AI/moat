
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDataSubjectRequestComponent } from './index.component';
import { DataSubjectRequestService } from '../../../services/DataSubjectRequest.service';

describe('IndexDataSubjectRequestComponent', () => {
  let component: IndexDataSubjectRequestComponent;
  let fixture: ComponentFixture<IndexDataSubjectRequestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDataSubjectRequestComponent
      ],
      providers: [
        DataSubjectRequestService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDataSubjectRequestComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});