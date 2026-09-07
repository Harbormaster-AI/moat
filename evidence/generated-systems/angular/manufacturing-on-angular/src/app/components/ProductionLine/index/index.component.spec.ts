
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProductionLineComponent } from './index.component';
import { ProductionLineService } from '../../../services/ProductionLine.service';

describe('IndexProductionLineComponent', () => {
  let component: IndexProductionLineComponent;
  let fixture: ComponentFixture<IndexProductionLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProductionLineComponent
      ],
      providers: [
        ProductionLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProductionLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});