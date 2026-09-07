
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexQuarantineComponent } from './index.component';
import { QuarantineService } from '../../../services/Quarantine.service';

describe('IndexQuarantineComponent', () => {
  let component: IndexQuarantineComponent;
  let fixture: ComponentFixture<IndexQuarantineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexQuarantineComponent
      ],
      providers: [
        QuarantineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexQuarantineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});