
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexEvidenceComponent } from './index.component';
import { EvidenceService } from '../../../services/Evidence.service';

describe('IndexEvidenceComponent', () => {
  let component: IndexEvidenceComponent;
  let fixture: ComponentFixture<IndexEvidenceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexEvidenceComponent
      ],
      providers: [
        EvidenceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexEvidenceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});